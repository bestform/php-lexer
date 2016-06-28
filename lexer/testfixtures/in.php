<?php

use Foo\Bar;

class Baz extends \Base implements \BaseInterface {

  /**
   * @param string $a
   * @param Type $b
   */
  public function fun1($a, Type $b) {
    // comment
    $s = $a;

    return true;
  }

}

/*===EXPECTED===
<?php
USE
NAME
;
CLASS
NAME
EXTENDS
NAME
IMPLEMENTS
NAME
{
DOCCOMMENT
PUBLIC
FUNCTION
NAME
(
VAR
,
NAME
VAR
)
{
COMMENT
VAR
=
VAR
;
RETURN
TRUE
;
}
}/*===EXPECTED===
