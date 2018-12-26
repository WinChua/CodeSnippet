#include <stdio.h>
#include <dirent.h>
#include <stdlib.h>


void show_dir(char * dirname) {
    DIR * dir = opendir(dirname);
    if( NULL == dir ) {
        perror(dirname);
        //exit(-1);
        return;
    }
    struct dirent *current = NULL;
    printf("%s\n", dirname);
    while((current = readdir(dir)) != NULL) {
        printf("\t%s\n", current->d_name);
    }
    closedir(dir);
    return;

}

int main(int argc, char * argv[]) {
    int i=0;
    for(i = 1; i < argc; ++i) {
        show_dir(argv[i]);
    }
    return 0;
}
