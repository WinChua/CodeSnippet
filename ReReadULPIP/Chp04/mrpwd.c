#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <dirent.h>

void get_inode(const char * filename, ino_t * res) {
    struct stat buf;
    stat(filename, &buf);
    *res = buf.st_ino;
    return;
}

char *inode_print_name(ino_t current) {
    DIR * p = opendir("..");
    struct dirent* d;
    while((d = readdir(p)) != NULL) {
        if(d->d_ino == current) break;
    }
    return d->d_name;
}

void get_parent(ino_t *sub) {
    ino_t parent_inode;
    get_inode("..", &parent_inode);
    if(parent_inode == *sub) {
        printf("/");
        return;
    }
    char *name = inode_print_name(*sub);
    chdir("..");
    get_parent(&parent_inode);
    printf("%s/", name);
}

void get_path() {
    ino_t sub;
    get_inode(".", &sub);
    get_parent(&sub);
    printf("\n");

}

int main(int argc, char * argv[]) {
    get_path();
    
}
