# Subtitles
A tool that extracts audio, trancscibes it, and writes it into subtitles from a video!
# Setup
1. Create either one or multiple accounts on https://www.assemblyai.com/ ( depending on how many API keys you are wanting to use since they have limits ) 
2. Get the API key from your account(s) and place them in the Config.json file.
3. Find audio from an external source and download it ( has to be a .mp4 file ) 
4. Place your downloaded audio in ./Assets/Videos/Input
5. Run the program with 2 arguments, the first is the Input file name, lets say you downloaded and put a file called audio.mp4 in the input folder, your first argument would be audio.mp4. The second argument is the Output file name, this can be whatever you like just make sure that the format you use for output is .m4a. It will then output your subtitles to ./Assets/Subtitles with a UUID name, it will tell you this name upon program completion.
6. Enjoy the tool.
