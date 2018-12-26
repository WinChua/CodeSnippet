#include <dirent.h>
#include <unistd.h>
#include <sys/stat.h>
#include <sys/types.h>

void get_inode(char *filename, ino_t * res) {
    struct stat buf;
    stat(filename, &buf);
    *res = buf.st_ino;
    return;
}

int main(int argc, char * argv[]) {
    ino_t current;
    get_inode(".", &current);
    
    ino_t parent;
    get_inode("..", &parent);
    while(parent != current) {
        DIR * p = opendir("..");
        struct dirent *s;
        while((s = readdir(p)) != NULL) {
            if(s->d_ino == current) {
                break;
            }
        }
        printf("%s/", s->d_name);
        current = parent;
        chdir("..");
        get_inode("..", &parent);
    }
}
