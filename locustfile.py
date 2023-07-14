from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def req_pq(self):
        self.client.post(
            '/auth/req_pq',
            json={
                     "nonce": "alpha",
                     "message_id": 7
            }
        )

    @task
    def get_users(self):
        self.client.post(
            '/biz/get_users',
            json={
                     "user_id": 0,
                     "message_id": 7,
                     "auth_key": 12
            }
        )
