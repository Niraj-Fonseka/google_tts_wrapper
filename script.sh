#!/bin/bash          
echo Hello World  

base64 dailyReport.txt -D > dest_audio_file.mp3

play dest_audio_file.mp3