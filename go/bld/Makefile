SRC_DIR := $(SOURCE_PATH)/src
OBJ_DIR := $(BUILD_DIR)/obj
INSTALL_DIR := $(BUILD_DIR)/install

.PHONY: all
all: install

GOFLAGS := $(GOFLAGS)

$(OBJ_DIR)/%.o : $(SRC_DIR)/%.go
	mkdir -p $(@D)
	$(GOC) -c $(GOFLAGS) $< -o $@

SRCs := 
SRCs += exr.go

SRCS := $(addprefix $(SRC_DIR)/,$(SRCs))

OBJs := $(SRCs:.go=.o)
OBJS := $(addprefix $(OBJ_DIR)/,$(OBJs))

$(OBJ_DIR)/exr : $(OBJS)
	# Link all the object files.
	$(GOC) $(GOFLAGS) $^ -o $@

$(OBJ_DIR):
	mkdir -p $@

$(INSTALL_DIR):
	mkdir -p $@

.PHONY: install
install: $(INSTALL_DIR) $(OBJ_DIR)/exr
	cp $(OBJ_DIR)/exr $(INSTALL_DIR)/.
	cp -r $(SRC_DIR)/html_templates $(INSTALL_DIR)/.

.PHONY: clean
clean:
	rm -rf $(OBJ_DIR)
	rm -rf $(INSTALL_DIR)


include $(SOURCE_PATH)/bld/doc.makefile
