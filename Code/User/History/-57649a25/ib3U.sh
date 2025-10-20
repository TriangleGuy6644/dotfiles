#!/bin/bash

set -euo pipefail
shopt -s nocasematch

ASCII=$(cat << "EOF"
 ███▄ ▄███▓ ▄▄▄        ▄████  ███▄    █  ██▓▄▄▄█████▓ █    ██ ▓█████▄ ▓█████ 
▓██▒▀█▀ ██▒▒████▄     ██▒ ▀█▒ ██ ▀█   █ ▓██▒▓  ██▒ ▓▒ ██  ▓██▒▒██▀ ██▌▓█   ▀ 
▓██    ▓██░▒██  ▀█▄  ▒██░▄▄▄░▓██  ▀█ ██▒▒██▒▒ ▓██░ ▒░▓██  ▒██░░██   █▌▒███   
▒██    ▒██ ░██▄▄▄▄██ ░▓█  ██▓▓██▒  ▐▌██▒░██░░ ▓██▓ ░ ▓▓█  ░██░░▓█▄   ▌▒▓█  ▄ 
▒██▒   ░██▒ ▓█   ▓██▒░▒▓███▀▒▒██░   ▓██░░██░  ▒██▒ ░ ▒▒█████▓ ░▒████▓ ░▒████▒
░ ▒░   ░  ░ ▒▒   ▓▒█░ ░▒   ▒ ░ ▒░   ▒ ▒ ░▓    ▒ ░░   ░▒▓▒ ▒ ▒  ▒▒▓  ▒ ░░ ▒░ ░
░  ░      ░  ▒   ▒▒ ░  ░   ░ ░ ░░   ░ ▒░ ▒ ░    ░    ░░▒░ ░ ░  ░ ▒  ▒  ░ ░  ░
░      ░     ░   ▒   ░ ░   ░    ░   ░ ░  ▒ ░  ░       ░░░ ░ ░  ░ ░  ░    ░   
       ░         ░  ░      ░          ░  ░              ░        ░       ░  ░
EOF
)

DESKTOPS=("KDE Plasma" "GNOME" "Budgie" "i3" "XFCE" "Cinnamon" "MATE" "LXQt" "Deepin" "Hyprland" "Wayfire" "Sway")
selected=0

draw_menu() {
    clear
    echo "$ASCII"
    echo
    echo "Select your Desktop Environment using arrow keys and Enter:"
    for i in "${!DESKTOPS[@]}"; do
        if [[ $i -eq $selected ]]; then
            printf "> \e[7m%s\e[0m\n" "${DESKTOPS[i]}"
        else
            printf "  %s\n" "${DESKTOPS[i]}"
        fi
    done
}

# Arrow key menu
while true; do
    draw_menu
    IFS= read -rsn1 key
    if [[ $key == $'\x1b' ]]; then
        IFS= read -rsn2 -t 0.1 key2
        key+="$key2"
        case "$key" in
            $'\x1b[A') ((selected--)); ((selected<0)) && selected=$((${#DESKTOPS[@]}-1)) ;;
            $'\x1b[B') ((selected++)); ((selected>=${#DESKTOPS[@]})) && selected=0 ;;
        esac
    elif [[ $key == "" ]]; then
        break
    fi
done

CHOICE="${DESKTOPS[selected]}"
clear
echo "$ASCII"
echo
echo "Installing: $CHOICE"

# Desktop installation commands
case "$CHOICE" in
    "KDE Plasma")
        sudo pacman -S --needed plasma kde-applications sddm --noconfirm
        sudo systemctl enable sddm --now
        ;;
    "GNOME")
        sudo pacman -S --needed gnome gdm --noconfirm
        sudo systemctl enable gdm --now
        ;;
    "Budgie")
        sudo pacman -S --needed budgie-desktop gdm --noconfirm
        sudo systemctl enable gdm --now
        ;;
    "i3")
        sudo pacman -S --needed i3-wm i3status dmenu --noconfirm
        ;;
    "XFCE")
        sudo pacman -S --needed xfce4 xfce4-goodies lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "Cinnamon")
        sudo pacman -S --needed cinnamon lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "MATE")
        sudo pacman -S --needed mate mate-extra lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "LXQt")
        sudo pacman -S --needed lxqt sddm --noconfirm
        sudo systemctl enable sddm --now
        ;;
    "Deepin")
        sudo pacman -S --needed deepin deepin-extra lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "Hyprland")
        sudo pacman -S --needed hyprland waybar --noconfirm
        ;;
    "Wayfire")
        sudo pacman -S --needed wayfire waybar --noconfirm
        ;;
    "Sway")
        sudo pacman -S --needed sway waybar --noconfirm
        ;;
esac

echo "$CHOICE installation complete!"
