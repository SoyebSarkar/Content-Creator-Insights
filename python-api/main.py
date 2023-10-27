from flask import Flask, request
import os
from googleapiclient.discovery import build

app = Flask(__name__)
youtube = build('youtube', 'v3', developerKey='AIzaSyB7VjizZKnfAVh5z49B0u26r7GqV5t6Ubg')

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
        responseData = {
            "video_id": video['id'],
            "video_title": video['snippet']['title'],
            "video_thumbnail": video['snippet']['thumbnails']['standard']['url'],
            "video_views": video['statistics']['viewCount'],
            "video_likes": video['statistics']['likeCount'],
            "video_comments": video['statistics']['commentCount']
        }
        responseObj.append(responseData)
    print(responseObj)
    return responseObj


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
