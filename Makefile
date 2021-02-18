

################################################################################################
# update requirements:
################################################################################################

# update package:
.PHONY:
go.mod.tidy:
	go mod tidy -v

go.require.add:
	go get -u github.com/tal-tech/go-zero
