#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <ctype.h>
#include <time.h>
#include <string.h>

int main(int argc, char *argv[]) {
    int c;

    opterr = 0;

    // get current date
    time_t t = time(NULL);
    struct tm date = *localtime(&t);

    const char *locations[] = {"mensa-garching","mensa-arcisstrasse","stubistro-grosshadern"};
    const int locations_len = sizeof(locations) / sizeof(locations[0]);
    char *location = NULL;

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
            // check if the first non-option arg is a location name. Any other non-option args will be ignored
            if (location == NULL) {
                int i;
                // check if location name is in array of available locations
                for (i = 0; i < locations_len; i++) {
                    if (strcmp(locations[i], argv[optind]) == 0) {
                        location = argv[optind];
                        break;
                    }
                }
                // if the passed location could'nt be found and therefore the variable isn't set return an error
                if (location == NULL) {
                    fprintf(stderr, "Location '%s' not found.\n", argv[optind]);
                    return 0;
                }
            }

            optind++;
        }
    }

    char date_str[80];
    strftime(date_str,80,"%Y-%m-%d", &date);
    printf("Menu for '%s' on '%s':\n", location, date_str);

    return 0;
}