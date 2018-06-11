#compdef dt

_arguments \
  '1: :->level1' \
  '2: :->level2' \
  '3: :_files'
case $state in
  level1)
    case $words[1] in
      dt)
        _arguments '1: :(base64 completion date hash help html jwt random server uri uuid version)'
      ;;
      *)
        _arguments '*: :_files'
      ;;
    esac
  ;;
  level2)
    case $words[2] in
      base64)
        _arguments '2: :(decode encode)'
      ;;
      uri)
        _arguments '2: :(decode encode)'
      ;;
      *)
        _arguments '*: :_files'
      ;;
    esac
  ;;
  *)
    _arguments '*: :_files'
  ;;
esac
