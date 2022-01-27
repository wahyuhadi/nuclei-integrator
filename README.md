# nuclei-integrator

### Scan your code with custom rules and integrate it with anything, anytime, anywhere


----

# How to install 
- add this to .bashrc 
```sh
export jira_user=<Jira email>
export jira_token=<jira token>
export jira_url=<Jira Url>
export elastic_url=<Elastic Url>
export elastic_user=<Elastic User>
export elastic_password=<Elastic password>
export elastic_index=<Elastic Index>
```

- and run unitest to makesure this code pass (dont install if found any fail in unitest)
```sh
source ~/.bashrc
make test 
```

- after run unitest 
```sh
make install 
```
- and running
```bash
nuclei -l target -silent -json | nuclei-integrator -jira // for jira
nuclei -l target -silent -json | nuclei-integrator -ep // for elastic push
```
# Will be integrate with:
- Elastic (done)


Push data to ELK
---
![sc](/img/nuclei.png)


Auto create issue jira
---
![sc](/img/jira.png)