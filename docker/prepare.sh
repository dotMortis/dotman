#!/bin/bash

pacman -Sy --noconfirm
pacman -Syu --noconfirm
pacman -S --needed go sudo nano --noconfirm

# Create a normal shell user named 'tu' (testuser)
useradd -m -s /bin/bash tu
usermod -aG wheel tu
chown -R tu:tu /workspace
# Set password for user 'tu' to 'tu'
echo "tu:tu" | chpasswd

# Uncomment the wheel group in sudoers
sed -i 's/# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/' /etc/sudoers
