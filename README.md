# NullDaddy-DDNS

 NullDaddy-DDNS is a weekend's long pet project attempting to provide a [poor-man's](http://www.urbandictionary.com/define.php?term=poor%20man%27s), [Dynamic-DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) using [ipify](https://www.ipify.org) and [GoDaddy](https://godaddy.com/)'s 
\- [Developer](https://developer.godaddy.com/) APIs.

## Usage

 NullDaddy-DDNS can be used as a CLI tool to issue a single update to a GoDaddy DNS record i.e.
<pre>
./nulldaddy-ddns --domain=[DOMAIN] --record=[DOMAIN_RECORD] --key=[DEVELOPER_KEY] --secret=[DEVELOPER_SECRET]
</pre> 
 
 NullDaddy-DDNS could also be used as a daemon to continuasly update the record in specified intervals by using
the "--daemon" flag.

  Please advise the "Help page" for more.

## Docker

 A docker image with NullDaddy-DDNS is provided at [docker hub](https://hub.docker.com/r/ipolyzos/nulldaddy-ddns).
 Environment variables could be used to configure all aspects of NullDaddy-DDS (please refer to help page) e.g.

<pre>
 docker run \
        -e GODADDY_DOMAIN=[DOMAIN_NAME] \
        -e GODADDY_RECORD=[SUBDOMAIN] \
        -e GODADDY_KEY=[GODADDY_DEVELOPER_KEY] \
        -e GODADDY_SERCRET=[GODADDY_DEVELOPER_SECRET] 
  nulldaddy-ddns:0.0.1
</pre>

## Help page

<pre>
NAME:
   nulldaddy-ddns - Poor man's DDNS!

USAGE:
   nulldaddy-ddns [global options] (mode)

VERSION:
   0.0.1

AUTHOR(S):
   Ioannis Polyzos <i.polyzos@null-box.org>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --key value, -k value           The developer's 'KEY'. [$GODADDY_KEY]
   --secret value, -s value        The developer's 'SECRET'. [$GODADDY_SERCRET]
   --domain value, --dn value      The 'DOMAIN' you would like to update. [$GODADDY_DOMAIN]
   --record value, -r value        The domain 'RECORD' you would like to update. [$GODADDY_RECORD]
   --record-type value, -t value   The 'TYPE' of domain record you would like to update. (default: "A") [$GODADDY_RECORD_TYPE]
   --record-ttl value, --tt value  The 'TTL' of domain record you would like to update. (default: 600) [$GODADDY_RECORD_TTL]
   --daemon, -d                    Run godaddy-dns as a'DAEMON'. [$NULLDADDY_DAEMON]
   --interval value, -i value      The 'INTERVAL' between update in sec. (default: 1800) [$NULLDADDY_INTERVAL]
   --help, -h                      show help
   --version, -v                   print the version
</pre>


## Developing NullDaddy-DDNS

See the [contribution guidelines](https://github.com/ipolyzos/nulldaddy-ddns/blob/master/CONTRIBUTING.md)