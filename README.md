# Daily_updates

Linux Commands and Tips
This document provides an overview of common Linux commands and related tips for working with files, directories, environment variables, and package management. The commands are organized into categories for easy reference.

Basic Commands
1. pwd - Present Working Directory
Description: Displays the current directory.
Example: pwd
2. cd - Change Directory
Description: Change the current directory.
Examples:
cd directory_name
cd .. - Move to the parent directory.
cd - - Return to the previous directory.
3. ls - List Directory Contents
Description: List the contents of a directory.
Examples:
ls
ls -l
...

Environment Variables
5. Environment Variables
PATH: Holds the executable file locations.
OLDPWD: Stores the previous directory.

echo $PATH
echo $OLDPWD


File Operations
4. cat - Concatenate and Display Files
Description: Display file contents and manipulate files.
Examples:
cat file1
cat file1 file2
cat file1 file2 > newFile
cat file1 >> file2
5. man - Manual Pages
Description: Access detailed information about commands.
Example: man command_name
...

Package Management
6. yum Package Manager (Red Hat based systems)
Description: Manage packages using Yellowdog Updater Modifier (yum).
Examples:
yum search package_name
yum install package_name
yum remove package_name


7. APT Package Manager (Debian/Ubuntu based systems)
Description: Manage packages using Advanced Package Tool (APT).
Examples:
sudo apt update
sudo apt install package_name
sudo apt remove package_name
