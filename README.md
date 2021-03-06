# **TuneIn**
TuneIn allows one to create a personal radio station from a repository of audio files. The radio station can be shared with just a web link and the audience requires no more than a modern browser.

After configuring the playlist and running the application, user can share the application link to her friends to listen the same radio channel.

![Tuning-image](https://github.com/gophergala2016/tunein/blob/master/screenshots/Tuning.png "Tuning")
![Playing-image](https://github.com/gophergala2016/tunein/blob/master/screenshots/Playing.png "Playing")


## Usage
- You can place your playlist in `music` folder of the source code.
- Build the code
- Run the application
- Open the browser and wait for the tuning time. Once tuning is done you can start listening to your favorite music.
- You can share the link to your friends so that they can also enjoy your playlist at the same time.


## Installation
- Install [`ffmpeg`](https://www.ffmpeg.org/)

  for Mac, you can use

  ```
  brew install ffmpeg
  ```
- Clone the TuneIn Project

  ```
  git clone https://github.com/gophergala2016/tunein.git
  ```
  
- Install golang websocket package

  ```
  go get golang.org/x/net/websocket
  ```
- Build the project using :

  ```
    $ go build
  ```
- Run the project (using default port 8080) :

  ```
    $ ./tunein
  ```
- You can specify your port using -p option :

  ```
    $ ./tunein -p 4000
  ```


## Idea
  This idea came to my mind based on the below scenarios.
  - In the absence of any good audio music player, one can listen to her favorite songs available to her computer.
  - One can listen to her favorite music stored in her computer even from her mobile or tablet.
  - One can share her playlist to her friends by sharing the application link.


## Troubleshoot
- After running the application, if there is any problem to listen the song, please cross check [`ffmpeg`](https://www.ffmpeg.org/) is installed properly or not.

## Note
- For better use, please choose smaller songs. So far I have tested with .mp3, .ogg audio formats.
- At the start, it may take some time for tuning. But once tuned in, songs will be played smoothly.

## Coming Soon
- The first cut version is simplistic. Stay tuned for enhanced capabilities like more supported formats, program scheduling, user interface enhancements etc.
- Suggestions and inputs are always welcome.
