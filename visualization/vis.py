import requests
import pandas as pd
import matplotlib.pyplot as plt
from datetime import datetime


apiEndpoint = "http://localhost:8080/"

def visAllTasks():

    # Make an API call
    response = requests.get(apiEndpoint + "tasks")

    # Check the response status code
    if response.status_code == 200:
        # Request was successful
        data = response.json()  # Parse the response JSON data
    else:
        # Request was not successful
        print("Request failed with status code:", response.status_code)

    createdAt = [datetime.strptime(task['CreatedAt'], '%Y-%m-%dT%H:%M:%S.%f%z').strftime('%Y-%m-%d') for task in data]
    
    task_counts = {}
    for date in createdAt:
        task_counts[date] = task_counts.get(date, 0) + 1

    dates = list(task_counts.keys())
    values = list(task_counts.values())
    
    fig = plt.figure(figsize = (10, 5))
    
    # creating the bar plot
    plt.bar(dates, values, color ='maroon',
            width = 0.4)
    
    plt.xlabel("Days")
    plt.ylabel("No. of tasks created")
    plt.title("Tasks per day graph")
    plt.show()

visAllTasks()