   #include <iostream>
   #include <boost/asio.hpp>
   
   class TCPServer {
   public:
       TCPServer(boost::asio::io_context& io_context, short port)
           : acceptor_(io_context, boost::asio::ip::tcp::endpoint(boost::asio::ip::tcp::v4(), port)) {
           start_accept();
       }
   private:
       void start_accept() {
           boost::asio::ip::tcp::socket socket(acceptor_.get_executor().context());
           acceptor_.async_accept(socket,
                                  [this](boost::asio::error_code ec) {
                                      if (!ec) {
                                          handle_connection(std::move(socket));
                                      }
                                      start_accept();
                                  });
       }
       void handle_connection(boost::asio::ip::tcp::socket socket) {
           try {
               boost::asio::streambuf buffer;
               boost::asio::read_until(socket, buffer, "\n");
               std::istream is(&buffer);
               std::string data;
               std::getline(is, data);
               std::cout << "Received: " << data << std::endl;
               std::string response = "Message received successfully";
               boost::asio::write(socket, boost::asio::buffer(response + "\n"));
           } catch (std::exception& e) {
               std::cerr << "Error: " << e.what() << std::endl;
           }
       }
       boost::asio::ip::tcp::acceptor acceptor_;
   };
   int main() {
       try {
           boost::asio::io_context io_context;
           TCPServer server(io_context, 8888);
           io_context.run();
       } catch (std::exception& e) {
           std::cerr << "Exception: " << e.what() << std::endl;
       }
       return 0;
   }