

configDirectory="$HOME/.config/gopsalm"

if [ -d "$configDirectory" ]; then
    rm -rf "$configDirectory"
fi

# Create the directory
mkdir -p "$configDirectory"


echo "======================================="
echo "Welcome, how are u? :D"
echo "======================================="
echo "Do you want everything in Spanish Or English?"
echo "Type 1 for Spanish 2 for English"
read choice
echo "======================================="

if [ "$choice" = "1" ]; then

    echo "Spanish selected"
    echo "Thanks, now I am going to install everything for u"
    cp -r DisplayEs "$configDirectory/Display"
elif [ "$choice" = "2" ]; then

    echo "English selected"
    echo "Thanks, now I am going to install everything for u"
    cp -r DisplayEn "$configDirectory/Display"
else
    echo "Invalid choice, going with default (English)"
    echo "Thanks, now I am going to install everything for u"
fi

go build ./gopsalm.go 
sudo mv gopsalm /bin/gopsalm

last_line=$(tail -n 1 "$HOME/.zshrc")

if [ "$last_line" != "gopsalm" ]; then
    echo "gopsalm">> "$HOME/.zshrc"
fi
