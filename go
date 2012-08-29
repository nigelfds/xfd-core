#!/bin/bash -e

# ensure RVM is installed
if [ ! -d "$HOME/.rvm" ]; then
	read -p "Could not find RVM ($HOME/.rvm). Install RVM from the interwebs? y/n : " -n 1 -r
	if [[ $REPLY =~ ^[Yy]$ ]]
	then
		echo -e "\n\ninstalling rvm...\n"
		curl -L get.rvm.io | bash -s stable    
	else
		echo -e "\n\nThis environment wrapper needs RVM. You can use Bundler and Rake to continue manually. \nExiting automation wrapper."  			
		exit 1
	fi
fi

source ~/.rvm/scripts/'rvm'
rvm rvmrc trust
rvm rvmrc load .

# install gems using bundler
gem list | grep bundler  || gem install bundler --version 1.0.21 --no-rdoc --no-ri
bundle check || bundle install

bundle exec rake $@
