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

    for ($a = 1; $a < 5; $a++) {
      // no op
    }

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
FOR
(
VAR
=
NUMBER
;
VAR
<
NUMBER
;
VAR
++
)
{
COMMENT
}
RETURN
TRUE
;
}
}/*===EXPECTED===
