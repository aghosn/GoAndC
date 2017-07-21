SUBDIRS = SharedDynamic  Static
CLEANDIRS = $(SUBDIRS:%=clean-%)

all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@

clean: $(CLEANDIRS)

.PHONY: all clean $(SUBDIRS) $(CLEANDIRS)
