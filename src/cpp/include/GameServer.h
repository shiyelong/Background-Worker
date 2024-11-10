#pragma once
#include <algorithm>
#include <iostream>
#include <sstream>
#include <cstring>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <thread>
#include <windows.h>
#define WIN32_LEAN_AND_MEAN

#pragma comment(lib,"ws2_32.lib")

class GameServer {
public:
    GameServer();

    virtual ~GameServer();

    bool Init();

    void Run() const;

    void Close() const;

    void SendMessage(std::string msg);

    void ReceiveMessage(std::string msg);

private:
    bool InitSocket();

    void CloseSocket();

    void SendSocket();

    void ReceiveSocket();

    void SendSocketData();

    void ReceiveSocketData();

    void SendMessageData();

    void ReceiveMessageData();

private:
    SOCKET sock;
    SOCKADDR_IN addr;
};
