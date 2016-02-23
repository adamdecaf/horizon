;; horizon emacs methods

(defun horizon/connect-postgres()
  (interactive)
  (mine/sql 'postgres "horizon" "e06b4ed2b382f68" "192.168.99.100" "horizon" "./emacs/"))

(provide 'horizon-utils)
