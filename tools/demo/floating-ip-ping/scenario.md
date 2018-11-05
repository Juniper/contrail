# To verify Floating IP ping

1. Create virtual machine with atomized Contrail,
2. Open three terminals,
3. Log them on the machine,
4. Create k8s Namespace using kubectl (You can use [example namespace](examples/namespace.json)),
5. Create Project in VNC API or use default-project or k8s-default,
6. Create [Network IPAM](examples/ipam.yml) in VNC API with ipam_subnet_method: "user-defined-subnet",
7. Create [Virtual Network](examples/vn.yml) in VNC API with address_allocation_mode: "user-defined-subnet-only",
   network_ipam_refs set to previously created Network IPAM with added subnet in attr field,
8. Create [Floating IP Pool](examples/f-ip-pool.yml) in VNC API as a child of previously created
   Virtual Network with floating_ip_pool_subnets. You can check ipam_subnets within
   Network IPAM reference in Virtual Network,
9. Create two pods on different terminals to ensure normal ping between them is
   working,
10. Create clean [Floating IP](examples/f-ip.yml) in VNC API as a child of Floating IP Pool,
11. Check IP address allocated for Floating IP. Try to ping it from any pod
    (it should fail),
12. Find Virtual Machine Interface connected with one of previously created pods,
13. Add Virtual Machine Interface Ref to Floating IP,
    (Find Virtual Machine Interface which contains phrase "busybox" in its fq_name, then copy
    that fq_name and place it into [Floating IP](examples/f-ip-ref.yml),
14. Try to ping once again (now it should succeed),
15. Enjoy your pinging!
