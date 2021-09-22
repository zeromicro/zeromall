
install:
	brew install go-task/tap/go-task

init:
	task install
	task init
	make tips

tips:
	echo "now, your can use [task xxx] to do everything!"
	echo "now, your can use [task xxx] to do everything!"
	echo "now, your can use [task xxx] to do everything!"


tidy:
	task tidy

run:
	task run

# dry test: (like nginx -t, don’t run, just test the configuration file)
task-dry:
	task --dry
	task --dry install
	task --dry init
	task --dry go:install
	task --dry go:init
