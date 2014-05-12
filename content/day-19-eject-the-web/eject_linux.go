package eject

/*
   #include <fcntl.h>
   #include <linux/cdrom.h>
   #include <sys/ioctl.h>
   #include <sys/stat.h>
   #include <sys/types.h>
   #include <unistd.h>
   #include <errno.h>

   static int
   _eject(int f) {
     int r = 0;
     int fd = open("/dev/cdrom", O_RDONLY | O_NONBLOCK);
     if (fd == -1) {
       r = errno;
     } else {
       int e = CDROMEJECT;
       if (f == 0) {
         if (ioctl(fd, CDROM_DRIVE_STATUS, 0) == CDS_TRAY_OPEN)
           e = CDROMCLOSETRAY;
       } else if (f == 1)
         e = CDROMEJECT;
       else
         e = CDROMCLOSETRAY;
       if (ioctl(fd, e, 0) < 0) {
         r = errno;
       }
       close(fd);
     }
     return r;
   }
*/
import "C"
import "errors"
import "syscall"

func Eject() error {
	if r := C._eject(0); r != 0 {
		return errors.New(syscall.Errno(r).Error())
	}
	return nil
}
