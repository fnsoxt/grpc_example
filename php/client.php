<?php
ini_set('display_errors', 1);
ini_set('memory_limit', '1024M');
error_reporting(-1);

require_once '/data/wwwroot/g_vendor/autoload.php';
require_once dirname(__FILE__).'/client/GPBMetadata/Pb/HelloGrpc.php';
require_once dirname(__FILE__).'/client/Service/GreeterClient.php';
require_once dirname(__FILE__).'/client/Service/HelloRequest.php';
require_once dirname(__FILE__).'/client/Service/HelloReply.php';

function greet($name)
{
    $client = new Service\GreeterClient('127.0.0.1:8080', [
        'credentials' => \Grpc\ChannelCredentials::createInsecure(),
    ]);
    $request = new Service\HelloRequest();
    $request->setName($name);
    list($reply, $status) = $client->SayHello($request)->wait();
    $message = $reply->getMessage();

    return $message;
}

$name = !empty($argv[1]) ? $argv[1] : 'world';
echo "the greet: ".greet($name)."\n";
