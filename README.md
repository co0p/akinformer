why don't you ASCK ?
====================

scraping the aerzte kammer job offer for fun and playing with google app engine

Idea:
-----

Schedule a repeating scraping task. If the scraper examines the html anf if it finds a new entry
on the job offer page, an email with the new offer(s) is being send.

* cron.yaml - registers a periodic call to the service