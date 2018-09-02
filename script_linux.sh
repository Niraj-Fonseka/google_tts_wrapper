#!/bin/bash          
echo Hello World  

base64 dailyReport.txt -d > dest_audio_file.mp3

aplay dest_audio_file.mp3