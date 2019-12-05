### Image-resizer

This is a simple batch image resizer written in Golang. It will take images from a directory, resize them, and save them in another directory.

This program runs on Linux, Windows and Mac. Downloads are available on the releases tab.

#### Usage

##### On Windows

These instructions are aimed at end-users that do not have advanced technical knowledge.

The first thing you need to do is to download the tool. The release page of the tool is here with links for all the platforms is here: https://github.com/DanielCalvo/image-resizer/releases

For Windows, the file that you want is image-resizer-windows-amd64.zip:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture1.JPG)

When you click this, you'll very likely receive an automatic warning from your browser that this file may be dangerous. You can ignore this warning and click on "Keep" on your downloads tab to save the file. The program is not malware (nor does it need administrator rights to run or install) but if you're particularly paranoid, you don't have to use it.

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture2.JPG)

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture3.JPG)

You then can then use either the zip tool that comes with Windows or winrar to decompress the archive which will be on your downloads folder. Once you're done with that and you enter the image-resizer-windows-amd64 directory on your downloads folder, you should see the program and the configuration file:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture4.JPG)

Now that you have the program, let's select the images that we want to resize and put them somewhere so that our program can resize them. For this example, I created two folders on my desktop, one named images and the other named resized:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture7.JPG)

I then went around on my photos folder and copied a bunch of pictures I wanted to resize into the images folder on my desktop:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture8.JPG)

The next step is to open the parameters.txt file that is on your downloads folder together with the executable program. On that file we'll tell the program where to find the images to resize, where to save the resized copies, and to which percentage we want the images resized to.

When you first open the file, it will look like this:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture6.JPG)

We need to put the path to the folders on your desktop there. The easiest way to get the folder path on Windows is to click on the address bar on the file explorer, and copy that into the file:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture5.jpg)

So the parameters.txt file will look like this in my case:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture10.JPG)

Save the file and with that we can run the program! Except that... remember that program has no graphical interface? We need to run it from the command line. This is a bit different from how you run most Windows programs, but is also easy. Go to your downloads folder again and then to the image resizer folder, and copy the folder address from the file explorer:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture11.jpg)

Copy that text, go on your start menu and just type "cmd" and press enter on Windows 10. This will open a command prompt. type cd and CTRL+V with the path you had copied before and press enter. The result should look like this:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture12.JPG)

And now if you type image_resizer.exe and press enter, the magic should start happening!

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture13.JPG)

Going back to the "resized" folder on your desktop, you should see two folders in there: One named 20 and the other named 50, as these were the ratios we specified for the program to resize.

You can choose as many resize ratios as you want. You can also put additional images to resize on your "to resize" directory, as the program will ignore the ones already resized. I ran the program again with these values on parameters.txt:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/Capture14.JPG)

And now I have thumbnail sized copies of my pictures in 5% and 10% of their original size:

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/bike5.jpg)

![](https://raw.githubusercontent.com/DanielCalvo/image-resizer/master/doc_images/bike10.jpg)

Bear in mind that for the sake of this example, I created a directory on my desktop and pasted the images in there that I wanted to resize, but this is not necessary. You can simply specify any already existing directory with your pictures as the srcDir on parameters.txt and image-resizer will read all .jpg images it finds in there, and save resized copies of them to dstDir.

The program does not alter the original images, it only saves additional copies with a smaller size.

##### On Linux

On Linux things are a bit more straightforward. I'll assume you're somewhat familiar with the command line. Launch a terminal and:

Create a directory for your image resizer and go into it:

```bash
mkdir image-resizer
cd image-resizer/
```

Download the zip file with the executable and config:
```bash
wget https://github.com/DanielCalvo/image-resizer/releases/download/v0.1/image-resizer-linux-amd64.zip
```

Unzip it:
```
$ unzip image-resizer-linux-amd64.zip
```

And you can use your favorite text editor (vim, nano, emacs, or any other graphical interface one) to edit the parameters.txt file to indicate where your images are located, and where you would like to save the resized copies. By default this will be:
```yaml
SrcDir: /tmp/img
DstDir: /tmp/resized
Ratio: [20, 50]
```

You can create these directories under /tmp/ and put the images there, or you can specify something like:
```yaml
SrcDir: /home/daniel/Pictures/mymotopictures
DstDir: /home/daniel/Pictures/resizedmotopictures
Ratio: [20, 50]
```

Running the program is just a matter of typing:
```bash
./image_resizer
```

##### On Mac
I have compiled the program for Mac and made it available here: https://github.com/DanielCalvo/image-resizer/releases (as image-resizer-mac-amd64.zip, should work on any mac made in the last 10 years)

I don't have access to a Mac, so I could not actually test this :(

But the steps should be nearly the same as the Linux ones:
- Download the correct zip file (the mac one)
- Unzip it
- Specify the filesystem path on parameters.txt for the images you want to resize, and the filesystem path for where you would like to store the resized copies
- Specify the resize ratios you want
- Open a terminal, go to where the executable is and run the program from the command line with ./image-resizer

##### Building on Linux

###### For windows:
```bash
GOOS=windows GOARCH=amd64 go build -o image_resizer.exe image_resizer.go
```

###### For Linux
```bash
go build image_resizer.go 
```

###### For Mac
```bash
GOOS=darwin GOARCH=amd64 go build image_resizer.go
```

Built with `go version go1.13.4 linux/amd64`

