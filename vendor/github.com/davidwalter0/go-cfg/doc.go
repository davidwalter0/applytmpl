// Copyright (c) 2017 David Walter
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:

/*

cfg extract and apply configuration from the
environment or command line flags

   - types are inferred from structure member types
   - environment variable APP prefix inferred from struct name
   - value assignment priority is from top to bottom
     - struct tag
     - environment
     - flag
   - names
     - environment variable: recursive struct member name with upper case
       camel case underscore separated
     - flag recursive struct member name with lower case camel case
       hyphenated separated
   - flags override environment variables, which in turn override struct
     tags, which override the type default

   ```
   type S struct {
     i int              `name:"AyeAye"`                      // env name S_AYE_AYE       flag -aye-aye
     f float64          `default:"2.71728"`                  // env name S_F             flag -f
     M map[int]float64 `name:"Map" default:"e:2.71,pi:3.14"` // env name APP_MAP           flag -map
     A []string         `default:"a,b,c"`                    // env name S_A             flag -a
     outer struct {
       i int                                                 // env name S_OUTER_I       flag -outer-i
       inner struct {
          i int                                              // env name S_OUTER_INNER_I flag -outer-inner-i
       }
     }
   }
   ```

   - Exporting APP_OVERRIDE_PREFIX set to a value will override the
     structure name for the environment variable prefix e.g.

   ```
   export APP_OVERRIDE_PREFIX=APP
   type S struct {
     i int             `name:"AyeAye"`                       // env name APP_AYE_AYE       flag -aye-aye
     f float64         `default:"2.71728"`                   // env name APP_F             flag -f
     M map[int]float64 `name:"Map" default:"e:2.71,pi:3.14"` // env name APP_MAP           flag -map
     A []string        `default:"a,b,c"`                     // env name APP_A             flag -a
     outer struct {
       i int                                                 // env name APP_OUTER_I       flag -outer-i
       inner struct {
          i int                                              // env name APP_OUTER_INNER_I flag -outer-inner-i
       }
     }
   }

*/

package cfg
