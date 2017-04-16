#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <ctype.h>
#include <time.h>

int main(int argc, char *argv[]) {
    int c;

    opterr = 0;

    // get current date
    time_t t = time(NULL);
    struct tm date = *localtime(&t);

    /*
     * concepts partially adopted from:
     * - http://stackoverflow.com/questions/18079340/using-getopt-in-c-with-non-option-arguments
     * - https://www.gnu.org/software/libc/manual/html_node/Example-of-Getopt.html
     */
    while (optind < argc) {
        if ((c = getopt(argc, argv, "d:")) != -1) {
            switch (c) {
                case 'd':
                    // check if date conforms to required format
                    if (strptime(optarg, "%Y-%m-%d", &date) == NULL) {
                        fprintf(stderr, "Format of date '%s' is wrong. Required format: yyyy-mm-dd\n", optarg);
                        return 1;
                    }
                    //printf("%d-%d-%d\n", date.tm_year+1900, date.tm_mon+1, date.tm_mday);
                    break;
                case '?':
                    if (optopt == 'd')
                        fprintf (stderr, "Option -%c requires an argument.\n", optopt);
                    else if (isprint (optopt))
                        fprintf (stderr, "Unknown option `-%c'.\n", optopt);
                    else
                        fprintf (stderr,
                                 "Unknown option character `\\x%x'.\n",
                                 optopt);
                    return 1;
                default:
                    abort();
            }
        } else {
            // Regular argument
            printf("arg: %s\n", argv[optind]);
            optind++;
        }
    }

    return 0;
}