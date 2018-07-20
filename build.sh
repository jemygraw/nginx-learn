# --with-http_realip_module
# http://nginx.org/en/docs/http/ngx_http_realip_module.html
auto/configure --prefix=/User/jemy/Temp/nginx --with-http_realip_module

# --with-http_addition_module
# http://nginx.org/en/docs/http/ngx_http_addition_module.html
auto/configure --prefix=/Users/jemy/Temp/nginx --with-http_addition_module

# cd .. && git clone https://github.com/openssl/openssl.git
# --with-http_ssl_module --with-openssl=../openssl
# http://nginx.org/en/docs/http/ngx_http_ssl_module.html
auto/configure --prefix=/Users/jemy/Temp/nginx --with-http_ssl_module --with-openssl=../openssl