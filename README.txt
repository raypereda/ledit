$./replace -h

Usage of ./replace:
Replace is a tool for replacing lines in text files.
Lines matching a regular expression are replaced with fixed text.
  -infiles string
    	the glob pattern of input text files
  -pattern string
    	line pattern for searching for
  -replace string
    	replace line for matched lines

Replace one line a time by executing replace once per line. If that edit is good, 
move the .edited to original filename. And then repeat for each line edit.


./replace -infiles="input1.rds" -pattern=ConnectString \
   -replace="      <ConnectString>Data Source=DbServerName101;Initial Catalog=DatabaseName201</ConnectString>"

mv input1.rds.edit input2.rds

./replace -infiles="input2.rds" -pattern=IntegratedSecurity \ 
   -replace="   <IntegratedSecurity>False</IntegratedSecurity>"

mv input2.rds.edited input3.rds

input3.rds has both lines edited. DbServerName100 changed to DbServerName101, 
DatabaseName200 changed to DatabaseName201, and the value for <IntegratedSecurity> was changed to False.

The OSX executable is cmd/replace/replace, and the Windows 64-bit executable is in cmd/replace.exe.
