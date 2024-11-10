#include "GameServer.h"

GameServer::GameServer():sock(INVALID_SOCKET),addr() {
    std::cerr<<"Creating GameServer"<<std::endl;
}

GameServer::~GameServer() {
    std::cerr<<"Destroying GameServer"<<std::endl;
}

bool GameServer::Init() {
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        return false;
    }
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET) {
        return false;
    }
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");
    addr.sin_port = htons(9999);
    if (connect(sock, (SOCKADDR*)&addr, sizeof(addr)) == SOCKET_ERROR) {
        return false;
    }
    return true;
}

void GameServer::Run() const {
    std::cerr<<"Running GameServer"<<std::endl;
        int ret = listen(sock, 5);
        if (ret == SOCKET_ERROR) {
            std::cerr<<"Failed to listen"<<std::endl;
            return;
        }
        std::cerr<<"Waiting for client connection"<<std::endl;
}

void GameServer::Close() const {
    std::cerr<<"Closing GameServer"<<std::endl;
    shutdown(sock, SD_BOTH);
    closesocket(sock);
    WSACleanup();
}


