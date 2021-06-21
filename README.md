## SocialOnline
This is an online social application where users can post/delete the pictures or videos and search the post by username or by post keywords.  

The following diagram shows the design architecture of the project.  
The project was implemented on Google Virtual Machine.  
Backend: Language: Go  

         Database: elasticsearch & GCS  
                   Elasticsearch can support various search function. The user imformation and corresponding post are stored in ES. 
                   GCS is also used for media files storage.

![image](https://user-images.githubusercontent.com/70457942/122697483-73909980-d213-11eb-8d32-71c861c7a1a2.png)
