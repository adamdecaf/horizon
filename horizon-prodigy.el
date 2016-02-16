;; horizon prodigy commands

(defun horizon/prodigy-service(name cmd args)
  (prodigy-define-service
    :name (concat "horizon - " name)
    :command cmd
    :args args
    :cwd "~/src/go/src/github.com/adamdecaf/horizon/"
    :tags '(horizon)
    :kill-signal 'sigkill))

(defun horizon/prodigy-shell-script(name script-name)
  (horizon/prodigy-service name "bash" (list (concat "./bin/" script-name))))

;; start postgres
(horizon/prodigy-service "start postgres" "docker-compose" (list "up" "-d" "postgres"))

;; go commands
(horizon/prodigy-service "build source" "go" (list "build" "-v" "."))
(horizon/prodigy-shell-script "update deps" "get-deps.sh")
(horizon/prodigy-shell-script "run tests" "run-tests.sh")
(horizon/prodigy-shell-script "start apps" "local-run.sh")

;; take backup / restore backup
(horizon/prodigy-service "take postgres backup" "" (list ""))
(horizon/prodigy-service "restore postgres backup" "" (list ""))

(provide 'horizon-prodigy)
