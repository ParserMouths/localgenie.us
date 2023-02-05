<img title="a title" alt="Alt text" src="./readme-img/pic.png">

# local-genie

Have you ever had a craving for pani-puri waale bhaiya who makes the best
pani-puri around your town but you find that he is not near his spot.
What can you do? You can't even go on Google Maps or some other place to find
information about stall's status.

Local-Genie has appeared to fulfill this request of yours. It can not only help
you find your favorite panipuri stall but also any local stall which is mobile.
Our Application(PWA) focuses on creating a platform for local-stall owners to get
discovered by their surrounding. We also provide an interface for users to discover
all the food-varieties that their locality has to offer.

Our app also provides you with the reviews on those stalls which will help you make a
better decision for your next place. Our idea can be extended to included all kinds of
local stall owners who rely on the crowd and are dynamic according to the needs of the
user.

## Features

### From users point of view 🌚

- User get to track their favourite stalls, so their cravings doesn't get ignored.
- Can get review from other users about stalls.
- Get notified when their favorite stalls are in nearby area.
- Get map with all the nearby stalls.

### From vendors point of view 🌝

- Vendors get better engagement and hence better revenues.
- Helps to improve vendors to make their business better from the users.
- Can send notifications to the the users who marked stalls as favourite.
- Can help identify areas of high revenue. i.e. the place where they can open their stall to earn more revenue.


## Tech Stack

- We have used **StoryBlok**'s headless CMS to deliver stories to our frontend.
- It drastically reduced speeds for fetching the data than fetching it from our servers.
- We have also used StoryBlok's asset store for storing and delivering all the images. 

### Frontend🖥️ 
- React for UI and components
- Storyblok for components related to the information of stalls

### Backend 🛠️
- Backend is microservice created in **Go** created with [fiber](https://gofiber.io/) framework.
- Architecture is based on Domain Driven Design, for scalability and cleaner repository design.
- PostgresSQL is used as database since the models were relational.
- Docker and Docker compose for developement and deployment environments.
<img width="1392" alt="ddd" src="https://user-images.githubusercontent.com/51414879/216795052-af2b1353-db6b-4a8a-9009-3bf779b5a708.png">


### Deployments🚀
- Azure for the Virtual Machine
- GoDaddy(Porkbun) for domain
- Github for code collaboration. 
- site: localgenie.us


