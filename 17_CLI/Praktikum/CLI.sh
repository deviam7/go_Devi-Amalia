# membuat folder utama
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"

# membuat file facebook.txt dan linkedin.txt dengan argumen kedua dan ketiga
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/about_me/personal"
echo "https://facebook.com/deviam" > "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/about_me/personal/facebook.txt"
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/about_me/professional"
echo "https://linkedin.com/in/deviam07" > "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/about_me/professional/linkedin.txt"

# membuat file list_of_my_friends.txt dengan command curl dari file yang diberikan
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_friends"
curl"https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_friends/list_of_my_friends.txt"

# membuat file about_this_laptop.txt dengan nama user dan uname -a
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info"
touch "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"
echo "User: $(whoami)" > "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"
echo "System Info: $(uname -a)" >> "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"

# membuat file internet_connection.txt dengan hasil ping ke google.com sebanyak 3 kali
mkdir -p "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info"
touch "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"
echo "User: $(whoami)" > "$1Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"
echo "System Info: $(uname -a)" >> "$1Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$1 Devi $(date '+%a %B %d:%H:%M:%S WIB %Y')"/internet_connection.txt"