EXECS := snake

.PHONY: all clean

all: $(EXECS)

$(EXECS):
	go build -o $@.exe ./cmd/$@

clean:
	$(RM) $(EXECS)
