// "web" is for REST binary, which handles API related things.
// "worker" is the actual bot or the autobase itself.

// You can deploy both in the same heroku app, but it's not recommended if you deploy on free dyno.
// Since free dyno will sleep app all process with contains "web" process if it's unused, hence
// the bot will also be put to sleep.

// If it's really necessary to use the API on heroku server, I'd recommend to put them in separate apps
// since you I assume you won't really need 100% uptime for the "web", but you need it for the "worker"
// (Heroku won't put app to sleep if it only contains "worker" process).


web: bin/rest
// worker: bin/worker