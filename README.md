why don't you let AK infomer do the work ?
===========================================

crawling the aerzte kammer job offer for fun and playing with google app engine

how it works
-----------
 
 * parse website extract job offers
 * calculate checksum and put new/updated offers in db
 * send emails to subscribers informing them about those offers
 * repeat every n hour

 development
 -----------

    cd $GOROOT/src/path/to/project
    ./dev_appserver.py .

deployment
----------

    gcloud app deploy

    