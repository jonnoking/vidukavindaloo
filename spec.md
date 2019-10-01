# Viduka Vindaloo

## Spec

- [x] Github repo
- [x] Go static website
- [ ] clean architecture - medium post
- [ ] ngrok localhost config
- [x] Deploy website to docker
- [x] Script docker deployment
- [x] letsencrypt https
- [ ] DNS subdomain hosted on docker
- [ ] Add logging to Go-www service
- [x] FPL - API working
- [ ] FPL - get snapshot of team
- [ ] FPL - get snapshot of each team in league
- [ ] FPL - scheduled event to capture each team
- [ ] FPL - get team per league
- [x] FPL - get transfers per team
- [ ] FPL - get fixtures
- [ ] FPL - get event - https://fantasy.premierleague.com/api/event/3/live/


## Fantasy Premier League

https://www.reddit.com/r/FantasyPL/comments/c64rrx/fpl_api_url_has_been_changed/

https://fantasy.premierleague.com/api/bootstrap-static/

To use my team api authentication is required (see link https://medium.com/@bram.vanherle1/fantasy-premier-league-api-authentication-guide-2f7aeb2382e4)

https://fantasy.premierleague.com/api/my-team/{team-id}/

https://fantasy.premierleague.com/api/entry/{team-id}/

Classic league

- League ids for classic and h2h can be found in entry api.

https://fantasy.premierleague.com/api/leagues-classic/{league-id}/standings/

H2H Leagues

https://fantasy.premierleague.com/api/leagues-h2h/{league-id}/standings/

Transfer api

https://fantasy.premierleague.com/api/entry/{team-id}/transfers-latest/

