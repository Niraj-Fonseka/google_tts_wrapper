# Google_TTS_Wrapper

An API that leverages google's text to speech API and generates audio files and play them.


![architecture](/docs/GoogleTTS.jpg)



### Routes 


```
/health   --> Application health check 
/generate_news_report --> queries the top news headliens , generates an audio file , plays it 
/news/:source -> JSON output of the news depending on the source ( refer to https://newsapi.org/ for more info)
```

### Scripts 

```
script_pi.sh  --> decodes the base64 byte string into a mp3 file and then plays it 
script_mac.sh  --> decodes the base64 byte string into a mp3 file and plays it 

the only difference between these two scripts are the different command line tools used by raspberry pi ( raspbian ) and macos 
```