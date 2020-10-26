protogen_dir_go=/tmp/protogen

# regenerate go protobuf code
protoc-go: protoc-go-clean
	mkdir -p ${protogen_dir_go}/github.com/alexrudd/
	ln -s $(shell pwd) ${protogen_dir_go}/github.com/alexrudd/lb-teams
	find -iname '*.proto' -exec \
		protoc --proto_path=. --go_out=${protogen_dir_go}/ {} \
		\;

# remove all go generated protobuf code
protoc-go-clean:
	rm -rf ${protogen_dir_go}
	find -iname '*.pb.go' -exec rm {} \;