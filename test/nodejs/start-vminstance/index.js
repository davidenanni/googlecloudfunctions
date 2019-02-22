var http = require('http');

var Compute = require('@google-cloud/compute');

var compute = Compute();

exports.startInstance = function startInstance(req, res) {

    var zone = compute.zone('europe-west1-b');

    var vm = zone.vm('test-cloudfunc-02');

    const config = {
    	os : 'debian',
    	http: true,
    	https: true,
    	machineType: 'f1-micro',
    	metadata: {
    		items: [
      			{
        			key: 'startup-script',
					value: '#!/bin/bash \n wget https://raw.github.com/davidenanni/googlecloudfunctions/master/testVM/montecarloHTTPS \n wget https://raw.github.com/davidenanni/golang_gcp/master/WebServer/server.crt \n wget https://raw.github.com/davidenanni/golang_gcp/master/WebServer/server.key \n chmod -R 777 montecarloHTTPS \n ./montecarloHTTPS '
				}]
		}		
    };

    vm.create(config);


    console.log(`\n VM created successfully`);

	res.status(200).send('Success start instance');
};