
<?php


$content=curl_post_ssl("https://test.888dc.space:8080/",null);

var_dump($content);



function curl_post_ssl($url, $vars, $second=30,$aHeader=array())
{
    $ch = curl_init();
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0); // 检查证书中是否设置域名，如果不想验证也可设为0  
curl_setopt($ch, CURLOPT_VERBOSE, '1'); //debug模式，方便出错调试  
curl_setopt($ch,CURLOPT_SSL_VERIFYPEER,false);
curl_setopt($ch,CURLOPT_SSLCERT,"/Users/xiafei/test/gotest/twoWaySSL/keys/client.crt");
curl_setopt($ch,CURLOPT_SSLKEY,"/Users/xiafei/test/gotest/twoWaySSL/keys/client.key");
curl_setopt($ch, CURLOPT_CAINFO, '/Users/xiafei/test/gotest/twoWaySSL/keys/ca.crt');
curl_setopt($ch,CURLOPT_URL,$url);
if( count($aHeader) >= 1 ){
    curl_setopt($ch, CURLOPT_HTTPHEADER, $aHeader);
}
curl_setopt($ch,CURLOPT_POST, 1);
curl_setopt($ch,CURLOPT_POSTFIELDS,$vars);
$data = curl_exec($ch);
curl_close($ch);
if($data)
    return $data;
else   
    return false;
}