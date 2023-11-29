import os
from flask import Flask, request, jsonify
import pandas as pd
from transformers import AutoTokenizer
from transformers import AutoModelForSequenceClassification
from scipy.special import softmax
from googleapiclient.discovery import build
from tqdm import tqdm
import time
import mysql.connector





app = Flask(__name__)
# key1 = "AIzaSyB7VjizZKnfAVh5z49B0u26r7GqV5t6Ubg"
key2 = "AIzaSyBz5T6PD9tN5SSTEHVNXUG2HH6VZN1iEss"
youtube = build('youtube', 'v3', developerKey=key2)
MODEL = f"cardiffnlp/twitter-roberta-base-sentiment"
tokenizer = AutoTokenizer.from_pretrained(MODEL)
model = AutoModelForSequenceClassification.from_pretrained(MODEL)


conn = mysql.connector.connect(
    host="content-insight-app.c8iayctp2gb1.us-east-1.rds.amazonaws.com",
    user="admin",
    password="Test1234",
    database="content-insight-app"
)
cursor = conn.cursor()
QueryInsertCommentTable = "INSERT INTO `16_comment_stat` (`channel_id`, `video_id`, `comment`, `username`, `positive`, `negetive`, `neutral`) VALUES (%s, %s, %s, %s, %s, %s, %s)" 
QueryDeleteCommentTable = "DELETE FROM `16_comment_stat` WHERE `channel_id` = %s AND `video_id` = %s"


@app.route('/list/YT', methods=['POST'])
def listYT():
    # Access the ytRequest object within the function

    data = request.get_json()
    channel_id = data["channelID"]
    apiKey = data["apiKey"]

    videos = []
    next_page_token = None

    while True:
        ytRequest = youtube.search().list(
            part='id',
            channelId=channel_id,
            maxResults=50,
            # order='date',
            pageToken=next_page_token,
        )
        response = ytRequest.execute()


        for item in response['items']:
            if item['id']['kind'] == 'youtube#video':
                video_id = item['id']['videoId']
                video_info = get_video_info(video_id)
                videos.append(video_info)


        next_page_token = response.get('nextPageToken')


        if not next_page_token:
            break
    responseObj = []
    for video in videos:
        print(video)
        responseData = {
            "video_id": video['id'],
            "video_title": video['snippet']['title'],
            "video_thumbnail": video['snippet']['thumbnails']['default']['url'],
            "video_views": video['statistics'].get('viewCount',0),
            "video_likes": video['statistics'].get('likeCount',0),
            "video_comments": video['statistics'].get('commentCount',0)
        }
        responseObj.append(responseData)
    return responseObj


@app.route('/YT/analyse/<channelId>/<videoId>', methods=['GET'])
def analyse(channelId, videoId):
    comments = get_video_comments(videoId)
    comments_df = pd.DataFrame(comments)
    cursor.execute(QueryDeleteCommentTable, (channelId, videoId))


    analyseObj = []
    for i, row in tqdm(comments_df.iterrows(), total=len(comments_df)):
        try:
            temp = {}
            text = row['comment']        
            roberta_result = polarity_scores_roberta(text)
            # temp["comment"] = text
            # temp["scores"] = roberta_result
            # temp["username"] =row['username']
            # temp["likeCount"] = row['likeCount']
            # temp["date"] = row['date']
            # temp["channelName"] = row['channelName']
            cursor.execute(QueryInsertCommentTable, (channelId, videoId, text, row['username'], roberta_result['roberta_pos'],roberta_result['roberta_neg'], roberta_result['roberta_neu']))

            conn.commit()

            # analyseObj.append(temp)
        except Exception as e:
            print("Unable to execute for comment--->",text)
            print("Error--->",e)
        

    return jsonify({"status":200})






def polarity_scores_roberta(example):
    encoded_text = tokenizer(example, return_tensors='pt')
    output = model(**encoded_text)
    scores = output[0][0].detach().numpy()
    scores = softmax(scores)
    scores_dict = {
        'roberta_neg' : str(scores[0]),
        'roberta_neu' : str(scores[1]),
        'roberta_pos' : str(scores[2])
    }
    return scores_dict


def get_video_comments(video_id):
    comments = []
    next_page_token = "QURTSl9pMmF3VmxLQU1EcUt2ZjNZcnZrN2VrRHp5OTQydUlpZ0UxR2pzR21FdnNxNHc1MHBOLUQ0bDJfV1JmQWFGaUtqRDRSQWdvS3BMOA=="

    cnt = 3800
    while True:
        print(next_page_token)

        request = youtube.commentThreads().list(
            part='snippet',
            videoId=video_id,
            maxResults=1000,  # Adjust the number of comments per request
            pageToken=next_page_token,
        )
        print(request.to_json())
        # time.sleep(5)

        response = request.execute()

        print(cnt)
        cnt+=100
        for item in response.get('items', []):
            commentDetails = {}
            commentDetails['comment'] = item['snippet']['topLevelComment']['snippet'].get('textOriginal', '')
            commentDetails['username'] = item['snippet']['topLevelComment']['snippet'].get('authorDisplayName', '')
            commentDetails['likeCount'] = item['snippet']['topLevelComment']['snippet'].get('likeCount', 0)
            commentDetails['date'] = item['snippet']['topLevelComment']['snippet'].get('publishedAt', '')
            commentDetails['channelName'] = item['snippet']['topLevelComment']['snippet'].get('authorChannelUrl', '')
            comments.append(commentDetails)

        next_page_token = response.get('nextPageToken', None)


        if not next_page_token:
            break


    return comments









def get_video_info(video_id):
    ytRequest = youtube.videos().list(
        part='snippet,statistics',
        id=video_id,
    )
    response = ytRequest.execute()
    video_info = response['items'][0]
    return video_info

# main driver function

if __name__ == '__main__':
    app.run()
