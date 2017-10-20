var allApps = [
   {
       "name": "leaderboard",
       "img": "",
       "version": "1.0",
       "description": "Leader Board - MeteorJS sample app",
       "appServiceTemplates":
       [
           {
           "name": "meteor/leaderboard",
           "id": 0,
           "img": "https://photos-6.dropbox.com/t/2/AACUs_EfHMSP1ssZS_07m0OnVidQAsP-sb-Xi5HEnW6NRw/12/61840786/png/32x32/1/_/1/2/service-meteor.png/EPjMsbUEGBsgAigC/bnJ-VFB8AIfY6nfNFYp-J19eFV_xuw6pSN8tsY9MHIo?size=2048x1536&size_mode=3",           
           "version": "latest",
           "psbType": "k8s",
           "dependencies":
           [
           {
           "type": "mongo-dsb",
           "name": "mongo",
           "description": "Mongo Store",
           "img": "https://photos-5.dropbox.com/t/2/AAA9OjbOLAkXuoKJ_1_QYWQyPHkCI4kLm6-EoYJDiU84PA/12/61840786/png/32x32/1/_/1/2/infrastructure-service-mongodb.png/EPjMsbUEGBsgAigC/WPYbRqL8jhtq9i8ThE8kJMPBLPriYzG0lYLITv909R4?size=2048x1536&size_mode=3"
           }
           ],
           "exposedPorts":
           [
           80
           ],
           "httpPort": 80
           }
       ],
       "entryPointServiceURI": "meteor/leaderboard",
       "entryPointServicePath": ""
   },
   {
       "name": "lets-chat",
       "img": "",
       "version": "1.0",
       "description": "Let's Chat",
       "appServiceTemplates":
       [
           {
           "name": "sdelements/lets-chat",
           "id": 1,
           "img": "https://photos-3.dropbox.com/t/2/AAAOCQP6Ua06sFgRNEXms5jdRV13PrtRK2_xPu7s6grCpQ/12/61840786/png/32x32/1/_/1/2/service-nodejs.png/EPjMsbUEGBsgAigC/XcQ_662gSNNst_0c-hff2fnwYAs3UqoxplR2z8Arkvo?size=2048x1536&size_mode=3",
           "version": "latest",
           "psbType": "k8s",
           "dependencies":
           [
           {
           "type": "mongo-dsb",
           "name": "mongo",
           "img": "https://photos-5.dropbox.com/t/2/AAA9OjbOLAkXuoKJ_1_QYWQyPHkCI4kLm6-EoYJDiU84PA/12/61840786/png/32x32/1/_/1/2/infrastructure-service-mongodb.png/EPjMsbUEGBsgAigC/WPYbRqL8jhtq9i8ThE8kJMPBLPriYzG0lYLITv909R4?size=2048x1536&size_mode=3",
           "description": "Mongo Store"
           }
           ],
           "exposedPorts":
           [
           8080,
           5222
           ],
           "httpPort": 8080
           }
       ],
       "entryPointServiceURI": "sdelements/lets-chat",
       "entryPointServicePath": "login"
   },
   {
       "name": "wordpress",
       "img": "",
       "version": "1.0",
       "description": "WordPress Content Manager",
       "appServiceTemplates":
       [
           {
           "name": "wordpress",
           "img": "https://photos-2.dropbox.com/t/2/AACgdY0tchSXO8VoQ-q_-K0W3nEO9bnu9IuBBpOlhps-VA/12/61840786/png/32x32/1/_/1/2/service-apache-webserver.png/EPjMsbUEGBsgAigC/_d_6OzPbM7NUmmcr30e5ILSoh2CprWq3lB4dVsZ79Ws?size=2048x1536&size_mode=3",
           "id": 2,
           "version": "1.0",
           "psbType": "k8s",
           "dependencies":
           [
           {
           "type": "mysql-dsb",
           "name": "management-db",
           "img": "https://photos-5.dropbox.com/t/2/AAAVGmvtCmzZOec0pMjU8vWO5RipSD7XqYCk6_gOGpBSjg/12/61840786/png/32x32/1/_/1/2/infrastructure-service-mysql.png/EPjMsbUEGBsgAigC/G1g6Epzml0PnrvjNOQrRKUeZxWqpukUVkaFeIYOWPTo?size=2048x1536&size_mode=3",
           "description": "Management Database"
           },
           {
           "type": "docker-volume-dsb",
           "name": "document-volume",
           "img": "https://photos-2.dropbox.com/t/2/AAB9q9Qys2yja6N8H85SHlj-lXFyWdJ9liMhBbx87xu61Q/12/61840786/png/32x32/1/_/1/2/infrastructure-service-docker-volume.png/EPjMsbUEGBsgAigC/UchSHuuBKFLO7kU6IzRjrTsNebmwWQcD1Ih1_63cDDo?size=2048x1536&size_mode=3",
           "description": "Document Volume"
           }
           ],
           "exposedPorts":
           [
           8080
           ],
           "httpPort": 8080
           }
       ],
       "entryPointServiceURI": "wordpress",
       "entryPointServicePath": "/"
   },
   {
       "name": "buckets",
       "version": "1.0",
       "description": "Buckets CMS",
       "appServiceTemplates":
       [
           {
               "name": "hbokh/docker-buckets-cms",
               "id": 3,
               "img": "",
               "version": "latest",
               "psbType": "Shpocker",
               "dependencies":
               [
                   {
                   "type": "mongo-dsb",
                   "name": "mongo",
                   "img": "https://photos-5.dropbox.com/t/2/AAA9OjbOLAkXuoKJ_1_QYWQyPHkCI4kLm6-EoYJDiU84PA/12/61840786/png/32x32/1/_/1/2/infrastructure-service-mongodb.png/EPjMsbUEGBsgAigC/WPYbRqL8jhtq9i8ThE8kJMPBLPriYzG0lYLITv909R4?size=2048x1536&size_mode=3",
                   "description": "Document Store"
                   }
               ],
               "exposedPorts":
               [
                 3000
               ],
               "httpPort": 3000
           }
       ],
       "entryPointServiceURI": "hbokh/docker-buckets-cms",
       "entryPointServicePath": "/"
   },
   {
       "name": "hackathon",
       "img": "",
       "version": "1.0",
       "description": "Hackathon Registration App",
       "appServiceTemplates":
       [
           {
           "name": "hackathon",
           "id": 4,
           "img": "https://photos-5.dropbox.com/t/2/AABV2anYR9NpO65ZqfT75JWmXp8LzlgH4YnxidRUcdHgxg/12/61840786/png/32x32/1/_/1/2/service-java.png/EPjMsbUEGBsgAigC/9ibAtbAxQ58PzKK-hZUskWtekC9srXO1ZJ2rPxt55aA?size=2048x1536&size_mode=3",
           "version": "1.0",
           "psbType": "ShpanPaaS",
           "dependencies":
           [
           {
           "type": "h2-dsb",
           "name": "hackathon-db",
           "img":"https://photos-4.dropbox.com/t/2/AABDJlLvQG5WFwwkbaRFQSSjdOp3QdxO8Wn0h31ADq2d1A/12/61840786/png/32x32/1/_/1/2/infrastructure-service-postgres.png/EPjMsbUEGBsgAigC/_Z6crieRNi-VJKODmBmy9C-BKMdl7kNKEDhCED0n65A?size=2048x1536&size_mode=3",
           "description": "Hackathon ideas DB"
           },
           {
           "type": "shpanblob-dsb",
           "name": "hack-docs",
           "img":"https://photos-6.dropbox.com/t/2/AAA_xeUMG-SoMHcGbjCwBEv9Mbg1AcocBQbNwgXt3c2C3g/12/61840786/png/32x32/1/_/1/2/infrastructure-service-s3.png/EPjMsbUEGBsgAigC/grwqHHqLZAYoTKlhDP77hSfLNuePpvkteY-26sEe2Is?size=2048x1536&size_mode=3",
           "description": "Hackathon document store"
           }
           ],
           "exposedPorts":
           [
           8080
           ],
           "httpPort": 8080
           }
       ],
       "entryPointServiceURI": "hackathon",
       "entryPointServicePath": "hackathon-api/html"
   },
    {
      "name": "Hackathon Pro",
      "img": "",
      "version": "1.0",
      "description": "Hackathon Registration App",
      "appServiceTemplates":
      [
        {
          "name": "hackathon",
          "id": 5,
          "img": "https://photos-2.dropbox.com/t/2/AACgdY0tchSXO8VoQ-q_-K0W3nEO9bnu9IuBBpOlhps-VA/12/61840786/png/32x32/1/_/1/2/service-apache-webserver.png/EPjMsbUEGBsgAigC/_d_6OzPbM7NUmmcr30e5ILSoh2CprWq3lB4dVsZ79Ws?size=2048x1536&size_mode=3",
          "version": "1.0",
          "psbType": "ShpanPaaS",
          "dependencies":
          [
            {
              "type": "h2-dsb",
              "name": "hackathon-db",
              "img":"https://photos-4.dropbox.com/t/2/AABDJlLvQG5WFwwkbaRFQSSjdOp3QdxO8Wn0h31ADq2d1A/12/61840786/png/32x32/1/_/1/2/infrastructure-service-postgres.png/EPjMsbUEGBsgAigC/_Z6crieRNi-VJKODmBmy9C-BKMdl7kNKEDhCED0n65A?size=2048x1536&size_mode=3",
              "description": "Hackathon ideas DB"
            },
            {
              "type": "shpanblob-dsb",
              "name": "hack-docs",
              "img":"https://photos-6.dropbox.com/t/2/AAA_xeUMG-SoMHcGbjCwBEv9Mbg1AcocBQbNwgXt3c2C3g/12/61840786/png/32x32/1/_/1/2/infrastructure-service-s3.png/EPjMsbUEGBsgAigC/grwqHHqLZAYoTKlhDP77hSfLNuePpvkteY-26sEe2Is?size=2048x1536&size_mode=3",
              "description": "Hackathon document store"
            }
          ],
          "exposedPorts":
          [
            8080
          ],
          "httpPort": 8080
        },
        {
          "name": "submission evaluator",
          "id": 6,
          "img": "https://photos-5.dropbox.com/t/2/AABV2anYR9NpO65ZqfT75JWmXp8LzlgH4YnxidRUcdHgxg/12/61840786/png/32x32/1/_/1/2/service-java.png/EPjMsbUEGBsgAigC/9ibAtbAxQ58PzKK-hZUskWtekC9srXO1ZJ2rPxt55aA?size=2048x1536&size_mode=3",
          "version": "0.1",
          "psbType": "ShpanPaaS",
          "dependencies":
          [
            {
              "type": "rabbitmq-dsb",
              "img": "https://photos-4.dropbox.com/t/2/AABx5PVVxWy-jIZgigJH4QH9z2X4stzv_XlquAUsVigDkg/12/61840786/png/32x32/1/_/1/2/infrastructure-service-rabbitmq.png/EPjMsbUEGBsgAigC/I8QVMNx141UgAlsCEA3SAYye3OlpRlvZlQevM67AYMw?size=2048x1536&size_mode=3",
              "name": "messaging",
              "description": "Hackathon message bus"
            }
          ],
          "exposedPorts":
          [
            8080
          ],
          "httpPort": 8080
        }
      ],
      "entryPointServiceURI": "hackathon",
      "entryPointServicePath": "hackathon-api/html"
    }
];

export default allApps;
