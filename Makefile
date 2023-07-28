
EXE_DIR=./bin

all:
	[ -d $(EXE_DIR) ] || mkdir -p $(EXE_DIR)
	go build -o $(EXE_DIR)/cli ./cmd/cli
	go build -o $(EXE_DIR)/app ./cmd/app

clean:
	$(RM) bin/
