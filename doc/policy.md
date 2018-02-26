# Introduction

RBAC provides access control at API (operation) and resource level. Previously with multi-tenancy, only resource level access control modeled after unix style permissions for user, group (role) and others (world) was available .

RBAC currently works in conjunction with Keystone relying on user credentials obtained from keystone from token present in the API request. Credentials include user, role, tenant and domain information.

API level access is controlled by list of rules. Attachment point for the rules is global-system-config, domain and project. Resource level access is controlled by permissions embedded in the object.

# API Level access control

If RBAC feature is turned on, API server requires a valid token to be present in the **X-Auth-Token** of incoming request. If token is missing or is invalid, HTTP error 401 will be returned. API server trades token for user credentials (role, domain, project etc) from keystone.

api-access-list object holds rules of the form:

    <object, field> => list of <role:CRUD>

object refers to an API resource such as network or subnet. Field refers to any property or reference within the resource. Field is optional in which case the CRUD operation refers to entire resource. Each rule also specifies list of roles and their corresponding permissions as a subset of CRUD operations.

For example, ACL object for a project might look like this:

    <virtual-network, network-policy> => admin:CRUD
    <virtual-network, network-ipam> => admin:CRUD
    <virtual-network, *>    => admin:CRUD, Development:CRUD

Thus admin and users with role <em>Development</em> can perform CRUD operations on network in a project, only admin can perform CRUD operations for policy and IPAM inside a network.

Role is Keystone role name. Field can be resource property or Reference. Field can be multi level, for example network.ipam.host-routes (in first release only one level is supported).

Rule set for validation is union of rules from api-access-list object attached to :
 - user Project
 - user domain
 - default-domain

It is possible for project or domain ACL object to be empty. Access is only granted if a rule in the combined Rule set allows access. There is no explicit deny rule.

ACL object can be shared within a domain. Thus multiple projects can point to same ACL object (such as default).

# Object level access control

perms2 property of an object allows fine grained access control per resource. It has the following fields:
 - Owner (tenant uuid)
 - share list (list of (tenant/domain UUID, Permissions) tuple)
 - globally shared flag (plus Permissions)

Owner field is populated at the time of creation as per the logic in section 'Object level RBAC ownership' below. Share list gets built as object is selected for sharing with other users. Sharing can be enabled at tenant or domain level.

Permission field has following meaning:
  - R (Read object)
  - W (Create/Update object)
  - X (Link or refer to object)

Access is allowed if:
   - user is owner and permissions allow (rwx) or
   - user tenant or domain in shared list and permissions allow or
   - world access is allowed
