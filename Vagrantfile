VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"

  (1..4).each do |i|
    config.vm.define "horz-node-#{i}" do |cfg|
      cfg.vm.network "private_network", ip: "192.168.150.1#{i}"
      cfg.vm.host_name = "horz-node-#{i}.horizon"
      cfg.vm.provider "virtualbox" do |v|
        v.customize ["modifyvm", :id, "--memory", "1024"]
      end
    end
  end
end
