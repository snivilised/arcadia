
function auto-check() {
  local owner=$(git config --get remote.origin.url | cut -d '/' -f 4)
  local repo=$(git rev-parse --show-toplevel | xargs basename)

  echo "---> 😎OWNER: $owner"
  echo "---> 🎀REPO: $repo"

  update-mod-file $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  update-source-id-in-root-cmd $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi
  
  update-arcadia-in-root-cmd $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  update-arcadia-in-taskfile $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  update-workflow-names $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  rename-templ-data-id $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  update-import-statements $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  update-readme $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  rename-language-files $repo $owner
  if [ $? -ne 0 ]; then
    return 1
  fi

  reset-version
  if [ $? -ne 0 ]; then
    return 1
  fi

  touch ./.env
  echo "✔️ done"
  return 0
}

# the sed -i option edits the file in place, overwriting the original file
#
function update-all-generic() {
  local repo=$1
  local owner=$2
  local from=$3
  local name=$4
  local target=$5
  local replacement=$6

  find $from -name "$name" -type f -exec sed -i "s/${target}/${replacement}/g" {} +

  if [ $? -ne 0 ]; then
    echo "!!! ⛔ Aborted! update-all-generic failed for $owner/$repo:"
    echo "  ---> file pattern: $name"
    echo "  --->         from: $from"
    echo "  --->       target: $target"
    echo "  --->  replacement: $replacement"
    return 1
  fi

  return 0
}

function update-mod-file() {
  local repo=$1
  local owner=$2
  local from=./
  local file_pattern=go.mod
  local target="module github.com\/snivilised\/arcadia"
  local replacement="module github.com\/$owner\/$repo"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function update-source-id-in-root-cmd() {
  local repo=$1
  local owner=$2
  local from=./src/app/command
  local file_pattern=root-cmd.go
  local target="github.com\/snivilised\/arcadia"
  local replacement="github.com\/$owner\/$repo"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function update-arcadia-in-root-cmd() {
  local repo=$1
  local owner=$2
  local from=./src/app/command
  local file_pattern=root-cmd.go
  local target=arcadia
  local replacement=$repo
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function update-arcadia-in-taskfile() {
  local repo=$1
  local owner=$2
  local from=./
  local file_pattern=Taskfile.yml
  local target=arcadia
  local replacement=$repo
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function update-workflow-names() {
  local repo=$1
  local owner=$2
  local from=.github/workflows
  local file_pattern=*.yml
  local target="name: Arcadia"
  local tc_repo="$(echo ${repo:0:1} | tr '[:lower:]' '[:upper:]')${repo:1}"
  local replacement="name: $tc_repo"

  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function rename-templ-data-id() {
  local repo=$1
  local owner=$2
  local from=./
  local file_pattern=*.go
  local target="arcadiaTemplData"
  local replacement="${repo}TemplData"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function update-readme() {
  local repo=$1
  local owner=$2
  local from=./
  local file_pattern=README.md
  local target="arcadia: "
  local replacement="${repo}: "

  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
  if [ $? -ne 0 ]; then
    return 1
  fi

  target="snivilised\/arcadia"
  replacement="$owner\/$repo"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
  if [ $? -ne 0 ]; then
    return 1
  fi

  target="Arcadia Continuous Integration"
  tc_repo="$(echo ${repo:0:1} | tr '[:lower:]' '[:upper:]')${repo:1}"
  replacement="$tc_repo Continuous Integration"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
  if [ $? -ne 0 ]; then
    return 1
  fi

  return 0
}

function update-import-statements() {
  local repo=$1
  local owner=$2
  local from=./
  local file_pattern=*.go
  local target="snivilised\/arcadia"
  local replacement="$owner\/$repo"
  update-all-generic $repo $owner $from $file_pattern "$target" "$replacement"
}

function rename-language-files() {
  local repo=$1
  find . -name 'arcadia*.json' -type f -print0 |
  while IFS= read -r -d '' file; do
    mv "$file" "$(dirname "$file")/$(basename "$file" | sed "s/^arcadia/$repo/")"
  done
  return $?
}

function reset-version() {
  echo "v0.1.0" > ./VERSION
  echo "0.1.0" > ./src/app/command/version.txt
  return 0
}
