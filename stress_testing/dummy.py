# So far, dummy.py is necessary when starting a master, because locust needs such a file.
# Don’t worry, dummy.py has nothing to do with your test.

from locust import HttpUser, TaskSet, task, between


def index(l):
    l.client.get("/")


def stats(l):
    l.client.get("/stats/requests")


class UserTasks(TaskSet):
    # one can specify tasks like this
    tasks = [index, stats]

    # but it might be convenient to use the @task decorator
    @task
    def page404(self):
        self.client.get("/does_not_exist")


class WebsiteUser(HttpUser):
    """
    User class that does requests to the locust web server running on localhost
    """

    host = "http://127.0.0.1:8089"
    wait_time = between(2, 5)
    tasks = [UserTasks]