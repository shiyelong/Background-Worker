#pragma once
#include<algorithm>
#include<iostream>
#include<sstream>
#include<cstring>
#include<winsock2.h>
#include<ws2tcpip.h>

class GameServer {
public:
    GameServer();

    virtual ~GameServer() = 0;

    bool Init();

    void Run();

    void Close();

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
