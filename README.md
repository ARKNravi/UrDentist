# UrDentistApp

Application that can help prevention caries

<u>**By Ravenclaw**</u>

## 🔰 About

Saat ini prevalensi karies gigi di Indonesia sangatlah tinggi, khususnya pada anak-anak usia 5-9 tahun yang memiliki persentase sebesar 90%. Padahal gangguan kesehatan gigi dapat berdampak signifikan terhadap kualitas hidup sang anak, seperti mengalami speech delay, kurangnya kemampuan sosial, tidak bisa hadir sekolah, dan aktivitas lainnya. Gangguan tersebut akan berdampak pada kualitas partisipasi generasi mendatang dalam pertumbuhan ekonomi nasional dari berbagai aspek. Besarnya dampak negatif yang dihasilkan, tidak disertai dengan kesadaran masyarakat akan pentingnya menjaga kesehatan gigi sejak dini. Menyadari permasalahan tersebut, ide kami hadir dengan goals utama untuk meningkatkan kesadaran para orang tua akan pentingnya menjaga kesehatan gigi anak sedini mungkin sebagai upaya preventif terhadap penyakit gigi dan mulut. Terdapat 3 langkah signifikan yang dilakukan oleh aplikasi kami, yaitu : 1. Membangun habit (habit tracker) yang mendukung kesehatan gigi, seperti kebiasaan menyikat gigi secara teratur dan tepat, menjaga konsumsi makanan, dan lain-lain. 2. Mendeteksi penyakit gigi sedini mungkin (screening kondisi gigi) untuk menentukan tindakan yang perlu diambil agar kondisinya tidak parah 3. Memperluas pengetahuan orang tua terkait kesehatan gigi sesuai keperluan anak (edukasi) sehingga dapat meningkatkan awarnesss terhadap penyakit gigi

## 🔧 Development

Here is a description of our apps development

### 📓 Tech Stack

List all the Tech Stack we use to build the system in this this project.

| No  | Tech                  | Details                                                           |
| --- | --------------------- | ----------------------------------------------------------------- |
| 1   | Flutter               | To build a beautiful and usefull FrontEnd App                     |
| 2   | Go                    | To build a fast and efficient Backend App                         |
| 3   | Google Cloud Platform | To provide all application needs related to server infrastructure |

### 🔩 Our Code Repository

- [FrontEnd](https://github.com/AhmadSultanMA/UrDentist)
- [BackEnd](https://github.com/ARKNravi/hackfest-be)

### 📷 API & Database Specification

You can just look at the `Docs` folder to check Our `API-Specification` and `Database-Specificasion`
Here is the link

1. [API-Specification](https://github.com/ARKNravi/hackfest-be/blob/main/api_urdentist.md)

### 📁 File Structure

Here is our File Structure

```
│   .env
│   .gcloudignore
│   .gitignore
│   app.yaml
│   go.mod
│   go.sum
│   main.go
│
├───.github
│   └───workflows
│           go.yml
│
├───controller
│       appointmentController.go
│       cameraController.go
│       dentistController.go
│       habitTrackerController.go
│       paymentController.go
│       profileController.go
│       questionController.go
│       userController.go
│
├───database
│       database.go
│
├───middleware
│       userMiddleware.go
│
├───model
│       appointmentModel.go
│       consultationModel.go
│       dentistModel.go
│       habitTrackerModel.go
│       paymentModel.go
│       profileModel.go
│       questionModel.go
│       ratingModel.go
│       serviceModel.go
│       tempuserModel.go
│       userModel.go
│
├───repository
│       appointmentRepository.go
│       dentistRepository.go
│       habitTrackerRepository.go
│       paymentRepository.go
│       profileRepository.go
│       questionRepository.go
│       userRepository.go
│
└───routes
        appointmentRoutes.go
        cameraRoutes.go
        dentistRoutes.go
        habitTrackerRoutes.go
        paymentRoutes.go
        profileRoutes.go
        questionRoutes.go
        userRoutes.go

```

| No  | File Name            | Details                                                                                                                  |
| --- | -------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| 1   | hackfest-be          | A Submodule from our Backend Repo you can just klik [here](https://github.com/ARKNravi/hackfest-be) to visit   |
| 2   | UrDentist            | A Submodule from our Frontend Repo you can just klik [here](https://github.com/AhmadSultanMA/UrDentist) to visit |

## 🌟 Credit

The Member of our team

1. Ananda Ravi Kuntadi
2. Ahmad Sultan 
3. Nabila Nafilia 
4. Noory Azyza

## 🔒License

© Ravenclaw - Hackfest by Google Indonesia 2024
