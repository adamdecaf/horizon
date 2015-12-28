;; horizon emacs methods

(defun horizon/connect-postgres()
  (interactive)
  (mine/sql 'postgres "horizon" "e06b4ed2b382f68" "192.168.59.103" "horizon" "."))

(provide 'horizon-utils)
